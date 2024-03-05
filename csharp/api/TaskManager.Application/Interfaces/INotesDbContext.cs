using Microsoft.EntityFrameworkCore;
using TaskManager.Model;

namespace TaskManager.Application.Interfaces;

public interface INotesDbContext
{
    public DbSet<Note> Notes { get; set; }
    public Task<int> SaveChangesAsync(CancellationToken cancellationToken);
}