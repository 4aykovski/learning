using Microsoft.EntityFrameworkCore;
using TaskManager.Application.Interfaces;
using TaskManager.Model;
using TaskManager.Service.EntityTypeConfiguration;

namespace TaskManager.Service;

public class NotesDbContext : DbContext, INotesDbContext
{
    public DbSet<Note> Notes { get; set; }

    public NotesDbContext(DbContextOptions<NotesDbContext> options) : base(options) {}

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.ApplyConfiguration(new NoteConfiguration());
        base.OnModelCreating(modelBuilder);
    }
}