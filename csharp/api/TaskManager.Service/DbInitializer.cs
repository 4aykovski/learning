namespace TaskManager.Service;

public class DbInitializer
{
    public static void Initialize(NotesDbContext context)
    {
        context.Database.EnsureCreated();
    }
}